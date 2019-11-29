function generate(goDirs) {
    const lines = [];
    goDirs.forEach(dir => lines.push(`go vet ./functions${dir}`));
    lines.push('');
    goDirs.forEach(dir => lines.push(`go test ./functions${dir}`));

    return lines;
}

module.exports = generate;