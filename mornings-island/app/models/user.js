import DS from 'ember-data';

export default DS.Model.extend({
  name: DS.attr('string'),
  email: DS.attr('string'),
  verified: DS.attr('boolean', {defaultValue: false}) 
  	//need to include mailchimp or sendgrid or so

});
