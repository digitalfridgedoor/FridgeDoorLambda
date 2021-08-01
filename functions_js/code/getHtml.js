const https = require("https");

function getHtml(url) {
    return new Promise<string>((resolve, reject) => {
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

exports.getHtml = getHtml;
