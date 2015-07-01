import DS from 'ember-data';

export default DS.Model.extend({
  body: DS.attr('string'),
  author: DS.hasMany('staff'), //includes photographers and such
  publishDate: DS.attr('date'),
  favorites: DS.attr('number'),
  timeToRead: DS.attr('number'),
  numComments: DS.attr('number'),
  summary: DS.attr('string'), 
});

//able to include photographers and other writers, etc..
