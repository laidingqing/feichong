'use strict';
define(['angular'], function(angular) {
  /* Services */
  angular.module('app.interceptor', [])
  .factory('httpRequestInterceptor', function () {
    return {
      request: function (config) {
        config.headers['Authorization'] = 'Bearer ' + localStorage.getItem("token");
        return config;
      }
    };
  })
})
