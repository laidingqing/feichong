'use strict';

var controllers = require('./controllers')
var services = require('./services')
var bootstrap = require('bootstrap')
var routerApp = angular.module('dbsApp', ['ui.router', 'ui.bootstrap', 'ui.bootstrap.tpls', 'app.services']);

routerApp.config(function ($stateProvider, $urlRouterProvider) {
    $urlRouterProvider.otherwise('/login');
    $stateProvider
        .state('login', { url: '/login', templateUrl: 'views/login.html?1', controller: controllers.HomeCtrl })
        .state('dashboard', { url: '/dashboard', templateUrl: 'views/dashboard.html?6', controller: controllers.DashboardCtrl })
        .state('dashboard.customers', { url: '/customers', templateUrl: 'views/customers.html?6', controller: controllers.CustomersCtrl })
        .state('dashboard.orders', { url: '/orders', templateUrl: 'views/orders.html?4', controller: controllers.OrdersCtrl })
        .state('dashboard.tracks', { url: '/tracks', templateUrl: 'views/tracks.html?2', controller: controllers.TracksCtrl })
});
