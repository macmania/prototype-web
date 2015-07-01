import Ember from 'ember';
import config from './config/environment';

var Router = Ember.Router.extend({
  location: config.locationType
});

Router.map(function() {
  this.route('login', {path: 'login'});
  this.route('article');
  this.route('staff');
  this.route('write_blog');
  this.route('home');
  this.route('register');
  this.route('user');
});

export default Router;
