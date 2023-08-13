function humanReadableFileSize(size, si = true) {
    // https://stackoverflow.com/a/20732091/8672525
    const exponent = si ? 1000 : 1024;
    const magnitude = size === 0 ? 0 : Math.floor(Math.log(size) / Math.log(exponent));
    const amount = Number((size / Math.pow(exponent, magnitude)).toFixed(2))
    if (magnitude === 0) {
        return `${amount} B`;
    }
    return `${amount} ${'?kMGT'[magnitude]}${si ? '' : 'i'}B`;
}

const compressionDiv = document.createElement("div");
compressionDiv.id = "compression";
document.body.append(compressionDiv);

function createCompressionElement(id, state) {
    const downloadDiv = document.createElement("div");

    const pathDiv = document.createElement("div");
    pathDiv.append(document.createTextNode(state.path));
    downloadDiv.append(pathDiv);

    const name = state.name ?? (/\/([^/]*)$/.exec(state.path)[1] || "root");

    const nameDiv = document.createElement("div");
    const nameSpan = document.createElement("span");
    nameSpan.append(document.createTextNode(name));
    const extensionSpan = document.createElement("span");
    extensionSpan.append(document.createTextNode("." + state.extension));
    nameDiv.append(nameSpan, extensionSpan);
    const renameButton = document.createElement("button");
    renameButton.append(document.createTextNode("Rename"));
    renameButton.addEventListener("click", () => {
        if (nameSpan.contentEditable === "true") {
            const downloads = JSON.parse(localStorage.getItem("downloads") ?? "{}");
            if (!downloads[id]) {
                return
            }
            downloads[id].name = nameSpan.innerText;
            renameButton.innerText = "Rename";
            nameSpan.contentEditable = "false";
            localStorage.setItem("downloads", JSON.stringify(downloads));
        } else {
            renameButton.innerText = "Apply";
            nameSpan.contentEditable = "true";
            setTimeout(() => {
                nameSpan.focus();
                const range = document.createRange();
                range.selectNodeContents(nameSpan);
                const selection = window.getSelection();
                selection.removeAllRanges();
                selection.addRange(range);
            });
        }
    });
    nameDiv.append(renameButton);
    downloadDiv.append(nameDiv);

    nameSpan.addEventListener("keypress", e => {
        if (e.key === "Enter") {
            e.preventDefault();
            renameButton.click();
        }
    });

    const progressDiv = document.createElement("div");
    const progress = `${(Math.min(Math.abs(state.progress), 1) * 100).toFixed(1)}`;
    progressDiv.classList.add("progress");
    if (state.progress < 0) {
        progressDiv.classList.add("failed");
    } else if (state.progress >= 1) {
        progressDiv.classList.add("finished");
    }
    const progressSpan = document.createElement("span");
    progressSpan.style.width = `${progress}%`;
    progressSpan.classList.add("bar");
    progressDiv.append(progressSpan);
    const progressValue = document.createElement("span");
    progressValue.style.zIndex = "2";
    progressValue.append(document.createTextNode(progress + "%"));
    progressDiv.append(progressSpan, progressValue);
    downloadDiv.append(progressDiv);

    if (state.size) {
        const sizeDiv = document.createElement("div");
        sizeDiv.append(document.createTextNode(humanReadableFileSize(state.size)));
        downloadDiv.append(sizeDiv);

        const dlButtonDiv = document.createElement("div");
        const downloadButton = document.createElement("button");
        downloadButton.append(document.createTextNode("Download"));
        downloadButton.addEventListener("click", () => {
            window.open(`/download/${id}/${nameSpan.innerText}.${state.extension}`, "_blank").focus();
        });
        dlButtonDiv.append(downloadButton);

        const shareButton = document.createElement("button");
        shareButton.append(document.createTextNode("Copy share link"));
        shareButton.addEventListener("click", () => {
            navigator.clipboard.writeText(`${window.location.origin}/download/${id}/${nameSpan.innerText}.${state.extension}`);
        });
        dlButtonDiv.append(shareButton);
        downloadDiv.append(dlButtonDiv);
    }

    downloadDiv.title = `Started: ${new Date(state.createdAt).toLocaleDateString("en-UK", {
        year: "numeric",
        month: "short",
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
    })}`;
    if (state.finishedAt) {
        downloadDiv.title += `\nFinished: ${new Date(state.finishedAt).toLocaleDateString("en-UK", {
            year: "numeric",
            month: "short",
            day: "2-digit",
            hour: "2-digit",
            minute: "2-digit",
            second: "2-digit",
        })}`;
    }

    downloadDiv.dataset.id = id;

    return downloadDiv;
}

async function getState(id) {
    const downloads = JSON.parse(localStorage.getItem("downloads") ?? "{}");
    const newSate = await fetch(`/status/${id}`).then(r => r.ok ? r.json() : null);

    if (newSate === null) {
        delete downloads[id];
        return;
    }

    downloads[id].progress = newSate.progress;
    if (newSate.finishedAt) {
        downloads[id].finishedAt = newSate.finishedAt;
    }
    if (newSate.downloadSize) {
        downloads[id].size = newSate.downloadSize;
    }

    localStorage.setItem("downloads", JSON.stringify(downloads));

    return downloads[id];
}

async function rebuildCompression() {
    while (compressionDiv.lastChild !== null) {
        compressionDiv.removeChild(compressionDiv.lastChild);
    }

    const downloads = JSON.parse(localStorage.getItem("downloads") ?? "{}");
    const downloadIds = Object.keys(downloads);

    const activeStates = await Promise.all(downloadIds.map(id =>
        fetch(`/status/${id}`).then(r => r.ok ? r.json() : null),
    ));

    for (let i = 0; i < downloadIds.length; i++) {
        const id = downloadIds[i];
        const state = activeStates[i];

        if (state === null) {
            delete downloads[id];
            continue;
        }

        downloads[id].progress = state.progress;
        if (state.finishedAt) {
            downloads[id].finishedAt = state.finishedAt;
        }
        if (state.downloadSize) {
            downloads[id].size = state.downloadSize;
        }
    }

    localStorage.setItem("downloads", JSON.stringify(downloads));

    console.log(downloads);

    for (const [id, state] of Object.entries(downloads)) {
        compressionDiv.append(createCompressionElement(id, state));
    }
}

rebuildCompression().catch(console.error);


async function watch(id) {
    let downloads = JSON.parse(localStorage.getItem("downloads") ?? "{}");

    let state = downloads[id];
    if (!state) {
        return
    }

    let compressionElem = document.querySelector(`[data-id="${id}"]`);
    if (!compressionElem) {
        const state = await getState(id);
        compressionElem = createCompressionElement(id, state);
        compressionDiv.append(compressionElem);
    }

    const update = async () => {
        const state = await getState(id);
        const newElem = createCompressionElement(id, state);
        compressionDiv.replaceChild(newElem, compressionElem);
        compressionElem = newElem;

        if (state.progress >= 1 || state.progress < 0) {
            return
        }
        setTimeout(update, 200);
    };
    setTimeout(update, 200);
}


for (const filesizeElem of document.getElementsByClassName("filesize")) {
    const byteCount = Number(/\d+/.exec(filesizeElem.innerText)[0]);
    if (Number.isNaN(byteCount)) {
        continue;
    }
    filesizeElem.innerText = humanReadableFileSize(byteCount);
}

for (const modifiedElem of document.getElementsByClassName("modified")) {
    const timestamp = Number(modifiedElem.innerText);
    const date = new Date(timestamp * 1000);
    modifiedElem.innerText = date.toLocaleDateString("en-UK", {
        year: "numeric",
        month: "short",
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit",
    });
}

for (const downloadElement of document.getElementsByClassName("download")) {
    const buttons = [...downloadElement.children];
    for (const button of buttons) {
        button.addEventListener("click", async () => {
            const state = await fetch(`/compress/${button.innerText}/${button.dataset.path}`, {
                method: "POST",
            }).then(res => res.json());
            console.log(state.createdAt)
            const downloads = JSON.parse(localStorage.getItem("downloads") ?? "{}");
            downloads[state.id] = {
                createdAt: state.createdAt,
                progress: state.progress,
                path: button.dataset.path,
                extension: button.innerText,
            };
            localStorage.setItem("downloads", JSON.stringify(downloads));
            watch(state.id).catch(console.error);
        });
    }
}