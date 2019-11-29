const fs = require('fs');
const gobuild = require('./gobuild');
const govetgotest = require('./govetgotest');
const templateyml = require('./templateyml');

function extension(file) {
    const match = /\.(\w+)$/.exec(file)
    if (match != null) {
        return match[1]
    }

    return null
}

const localBasePath = '../functions';
const functionsDir = fs.readdirSync(localBasePath);

const goDirs = []

functionsDir.forEach(dir => {
    console.log(dir)
    let relativePath = '/' + dir;
    let localPath = localBasePath + relativePath;
    const files = fs.readdirSync(localPath);
    let isGoDir = false;
    files.forEach(file => {
        const path = localPath + file;
        if (extension(path) == 'go') {
            isGoDir = true;
        }
    })
    if (isGoDir) {
        goDirs.push(relativePath);
    }
});

const gobuild_sh = gobuild(goDirs);
const govetgotest_sh = govetgotest(goDirs);
const template_yml = templateyml(goDirs);

fs.writeFileSync('../gobuild.sh', gobuild_sh.join('\n'));
fs.writeFileSync('../govet_gotest.sh', govetgotest_sh.join('\n'));
fs.writeFileSync('../template.yml', template_yml.join('\n'));