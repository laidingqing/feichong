'use strict';
define(['angular'], function(angular) {
  /* Services */
  angular.module('app.interceptor', [])
    .factory('httpRequestInterceptor', function($q, $rootScope) {
      return {
        request: function(config) {
          config.headers['Authorization'] = 'Bearer ' + localStorage.getItem("token");
          return config;
        },
        responseError: function(response) {
          if (response.status == 401) {
            $rootScope.$emit("httpRequestInterceptor", "notLogin", response);
          }
          return $q.reject(response);
        }
      };
    })
})
