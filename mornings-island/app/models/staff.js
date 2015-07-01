import DS from 'ember-data';

export default DS.Model.extend({
  firstName: DS.attr('string'),
  lastName: DS.attr('string'),
  
  /** able to validate the form ****@.com **/
  email: DS.attr('string'),

  miniBio: DS.attr('string'),
  twitter: DS.attr('string'),
  instagram: DS.attr('string'),
  pinterest: DS.attr('string'), //idk if this is a thing
  type: DS.attr('string'), //photographer, writer, etc...
  website: DS.attr('string'), 
  verified: DS.attr('boolean', { defaultValue: false }), 
  articlesWritten: DS.hasMany('article') 
});
