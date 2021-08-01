exports.handler = async function (event, context, callback) {
    console.log('Hello, world');
    console.log(event.Records[0].Sns);
};
