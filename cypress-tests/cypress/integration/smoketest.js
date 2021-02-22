/// <reference types="cypress" />

context('Smoke test', () => {

  it('can render the homepage', () => {
    cy.visit('/')
    cy.get('[data-cy=page-heading]').should('contain', 'Homepage')
  })

})