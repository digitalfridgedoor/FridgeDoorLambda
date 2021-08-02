import * as index from '../src/index';
import { expect } from 'chai';
import 'mocha';
import { createSNSMessage } from './test-data';

describe('Array', function () {
    describe('index', function () {
        it('Should get html', async () => {
            const message = createSNSMessage(JSON.stringify({ url: 'https://example.com/' }))
            await index.handler(message)
            // expect(x).to.equal('Hello World!');
        });
    });
});
