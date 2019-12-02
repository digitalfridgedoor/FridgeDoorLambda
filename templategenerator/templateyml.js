const yaml = require('js-yaml');
const fs = require('fs');

function readYmlTemplate(template) {
    let fileContents = fs.readFileSync(`templateyml/${template}.yml`, 'utf8');
    return fileContents;
}

function replaceProperty(original, property, replacement) {
    const regexp = new RegExp(`{{${property}}}`)
    return original.replace(regexp, replacement);
}

function generate(lambdaDefinitions) {
    const lines = [];
    const template = readYmlTemplate('base');
    lines.push(template);

    let functionTemplate = readYmlTemplate('function');
    lambdaDefinitions.forEach(definition => {
        const {
            localRelativePath,
            method,
            urlPath } = definition;

        let clean = localRelativePath
            .replace(/\//g, '') // replace all / with nothing
            .replace(/[\{\}]/g, ''); // replace brackets with nothing
        let template = functionTemplate;
        template = replaceProperty(template, 'Name', clean);
        template = replaceProperty(template, 'Handler', localRelativePath);
        template = replaceProperty(template, 'Path', urlPath);
        template = replaceProperty(template, 'Method', method);

        lines.push(template);
    });

    return lines;
}

module.exports = generate;