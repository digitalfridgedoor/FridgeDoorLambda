const yaml = require('js-yaml');
const fs = require('fs');

function readYmlTemplate(template) {
    let fileContents = fs.readFileSync(`templateyml/${template}.yml`, 'utf8');
    return fileContents;
}

function readFunctionTemplate(goDir) {
    try {
        let fileContents = fs.readFileSync(`${goDir}/template.yml`, 'utf8');
        return { exists: true, templateyml: yaml.safeLoadAll(fileContents) };
    } catch{
        return { exists: false };
    }
}

function replaceProperty(original, property, replacement) {
    const regexp = new RegExp(`{{${property}}}`)
    return original.replace(regexp, replacement);
}

function generate(goDirs) {
    const lines = [];
    const template = readYmlTemplate('base');
    lines.push(template);

    let functionTemplate = readYmlTemplate('function');
    goDirs.forEach(dir => {
        const { exists, templateyml } = readFunctionTemplate(`../functions${dir}`)
        if (exists) {
            let clean = dir.replace(/^\//, '');
            functionTemplate = replaceProperty(functionTemplate, 'Name', clean);
            functionTemplate = replaceProperty(functionTemplate, 'Handler', clean);
            functionTemplate = replaceProperty(functionTemplate, 'Path', templateyml[0]["Path"]);
            functionTemplate = replaceProperty(functionTemplate, 'Method', templateyml[0]["Method"]);

            lines.push(functionTemplate);
        }
    })

    return lines;
}

module.exports = generate;