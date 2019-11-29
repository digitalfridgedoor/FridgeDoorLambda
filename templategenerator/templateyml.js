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
            let template = functionTemplate;
            template = replaceProperty(template, 'Name', clean);
            template = replaceProperty(template, 'Handler', clean);
            template = replaceProperty(template, 'Path', templateyml[0]["Path"]);
            template = replaceProperty(template, 'Method', templateyml[0]["Method"]);

            lines.push(template);
        }
    })

    return lines;
}

module.exports = generate;