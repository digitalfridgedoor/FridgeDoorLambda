import { IngredientTagFinder } from '../../src/IngredientFinder/IngredientTagFinder'
import { expect } from 'chai';
import 'mocha';
import { TestCases } from './IngredientTagFinder.test-cases';

describe('IngredientTagFinder', () => {

    describe('find, no text', async () => {
        const html = `<div class="wrapper">
        <div class="ingredient" />
        <div class="ingredient" />
        <div class="ingredient" />
        <div>`

        const results = await IngredientTagFinder.find(html)

        it('Should find no results', async () => {
            expect(results.length).to.equal(0);
        });
    });

    describe('find, no children', async () => {
        const html = `<div class='wrapper'><div>`

        const results = await IngredientTagFinder.find(html)

        it('Should find no results', async () => {
            expect(results.length).to.equal(0);
        });
    });

    TestCases.forEach(({ html, expected, name }) => {
        describe(name, async () => {
            const sections = await IngredientTagFinder.find(html)

            it(name + ' Should have expected results', async () => {
                expect(sections.length).to.equal(expected.length);
                sections.forEach((section, idx) => {
                    const expectedSection = expected[idx]
                    expect(section.ingredients.length).to.equal(expectedSection.ingredients.length);
                    section.ingredients.forEach((ingredient, idx) => {
                        const expectedIngredient = expectedSection.ingredients[idx]
                        expect(ingredient).to.equal(expectedIngredient);
                    })
                })
            });
        });
    })
});
