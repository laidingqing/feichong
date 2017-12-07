'use strict';

var controllers = require('./controllers')
var services = require('./services')
var services = require('./filters')
var services = require('./directive')
var services = require('./interceptor')
var bootstrap = require('bootstrap')
var routerApp = angular.module('dbsApp', ['ui.router', 'ui.bootstrap', 'ui.bootstrap.tpls', '720kb.datepicker', 'app.services', 'app.filters', 'app.directive', 'app.interceptor']);

routerApp.config(function ($stateProvider, $urlRouterProvider, $httpProvider) {
    $httpProvider.interceptors.push('httpRequestInterceptor');
    $urlRouterProvider.otherwise('/login');
    $stateProvider
        .state('login', { url: '/login', templateUrl: 'views/login.html?6', controller: controllers.HomeCtrl })
        .state('dashboard', { url: '/dashboard', templateUrl: 'views/dashboard.html?8', controller: controllers.DashboardCtrl })
        .state('dashboard.customers', { url: '/customers', templateUrl: 'views/customers.html?6', controller: controllers.CustomersCtrl })
        .state('dashboard.orders', { url: '/orders', templateUrl: 'views/orders.html?8', controller: controllers.OrdersCtrl })
        .state('dashboard.tracks', { url: '/tracks', templateUrl: 'views/tracks.html?10', controller: controllers.TracksCtrl })
        .state('dashboard.business', { url: '/business/?:id&:orderNO', params: {'id': null, 'orderNO':null}, templateUrl: 'views/business.html?10', controller: controllers.BusinessCtrl })
        .state('dashboard.feedback', { url: '/feedback', templateUrl: 'views/feedback.html?6', controller: controllers.FeedbackCtrl })
});
