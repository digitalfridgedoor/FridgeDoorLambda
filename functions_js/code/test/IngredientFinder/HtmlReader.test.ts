import { HtmlReader } from '../../src/IngredientFinder/HtmlReader'
import { expect } from 'chai';
import 'mocha';
import { FindTestCase } from './IngredientTagFinder.test-cases';

describe('HtmlReader', () => {

    describe('readContents', async () => {
        it('simple example', async () => {
            const html = `<div class="wrapper">
        <div class="ingredient" />
        <div class="ingredient" />
        <div class="ingredient" />
        <div>`

            const section = await HtmlReader.readContents(html)

            console.log(JSON.stringify(section))

            expect(section.tag).to.equal('div');
            expect(section.classList.length).to.equal(1);
            expect(section.classList[0]).to.equal('wrapper');
            expect(section.children.length).to.equal(3);

            section.children.forEach(child => {
                expect(child.tag).to.equal('div')
                expect(child.classList.length).to.equal(1);
                expect(child.classList[0]).to.equal('ingredient');
                expect(child.children.length).to.equal(0);
            });
        });

        it('real html page', async () => {
            const html = FindTestCase('ratatouille').html;

            const section = await HtmlReader.readContents(html)
            console.log(section)

            expect(section.tag).to.equal('div');
            expect(section.classList.length).to.equal(1);
            expect(section.classList[0]).to.equal('wrapper');
            expect(section.children.length).to.equal(3);

            section.children.forEach(child => {
                expect(child.tag).to.equal('div')
                expect(child.classList.length).to.equal(1);
                expect(child.classList[0]).to.equal('ingredient');
                expect(child.children.length).to.equal(0);
            });
        });
    });
});
