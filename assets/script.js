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
