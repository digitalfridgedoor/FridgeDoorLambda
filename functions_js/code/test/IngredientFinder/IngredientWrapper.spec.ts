import { IngredientWrapperLocator } from '../../src/IngredientFinder/IngredientWrapperLocator';
import { HtmlSection } from '../../src/IngredientFinder/HtmlSection';
import { IngredientSection } from '../../src/IngredientFinder/Models';
import { expect } from 'chai';
import 'mocha';

describe('IngredientWrapper', () => {

  describe('find', async () => {

    it('Empty section has no results', async () => {

      const section: HtmlSection = {
        children: [],
        classList: [],
        tag: 'div'
      };

      const results = IngredientWrapperLocator.find(section);

      expect(Object.keys(results).length).to.equal(0);
    });

    it('Section with ingredient list', async () => {

      const section: HtmlSection = {
        children: [
          {
            children: [],
            tag: 'h2',
            text: 'Header'
          },
          {
            children: [
              {
                children: [],
                classList: ['ing'],
                tag: 'div',
                text: '1 pineapple'
              },
              {
                children: [],
                classList: ['ing'],
                tag: 'div',
                text: '1 ham'
              }
            ],
            classList: ['ing-wrapper'],
            tag: 'div'
          }
        ],
        classList: [],
        tag: 'div'
      }

      const results = IngredientWrapperLocator.find(section);

      it('Should find one ingredient list', async () => {
        expect(Object.keys(results).length).to.equal(1);
      });

      it('Ingredient list has correct items', async () => {
        expect(results[0].ingredients.length).to.equal(2);
        expect(results[0].ingredients[0]).to.equal('1 pineapple');
        expect(results[0].ingredients[1]).to.equal('1 ham');
      });
    });
  });
});
