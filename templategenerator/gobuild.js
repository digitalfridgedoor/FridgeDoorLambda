function generate(goDirs) {
    const lines = [];
    lines.push('mkdir bin');
    lines.push('');

    goDirs.forEach(dir => lines.push(`go build -o bin${dir} ./functions${dir}`))

    return lines;
}

module.exports = generate;