const https = require('https')
const fs = require('fs')
const rimraf = require('rimraf')

const testCasePath = './test/TestData'

function httpGet(url) {
    return new Promise((resolve, reject) => {
        https.get(url, (res => {
            if (res.statusCode === 200) {
                let data = '';
                res.on('data', d => data += d)
                res.on('end', () => resolve(data))
            } else {
                reject({ statusCode: res.statusCode, reason: 'invalid status' })
            }
        }))
    })
}

async function removeDir(dir) {
    return new Promise((resolve, reject) => {
        rimraf(dir, {}, err => {
            if (!err) {
                resolve()
            }
            reject(err)
        })
    })
}

async function run(url, file) {
    const html = await httpGet(url)
    fs.writeFileSync(testCasePath + '/' + file + '.html', html)
}

async function runAll() {
    await removeDir(testCasePath)

    fs.mkdirSync(testCasePath)

    await run('https://www.jamieoliver.com/recipes/cauliflower-recipes/spiced-whole-roast-cauli/', 'jamie_cauli')
    await run('https://www.recipetineats.com/chicken-sharwama-middle-eastern/', 'chicken_shwarma')
    await run('https://www.bbc.co.uk/food/recipes/red_pepper_and_aubergine_84745', 'shashuska_traybake')
    await run('https://www.delish.com/uk/cooking/recipes/a31011824/thai-green-curry/', 'thai_green_curry')
    await run('https://pinchofyum.com/creamy-thai-sweet-potato-curry', 'pot_curry')
}

runAll()
    .catch(x => console.log('error running', x))
