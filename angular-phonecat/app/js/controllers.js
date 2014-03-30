'use strict';

/* Controllers */

var phonecatControllers = angular.module('phonecatControllers', []);


phonecatControllers.controller('PhoneListCtrl', ['$scope', 'Phone',
  function($scope, Phone) {
    var q = Phone.query();
    //alert(JSON.stringify(q));
   //alert(q);
    //$scope.phones =  angular.fromJson(q);
    $scope.phones =  angular.fromJson(q);
    $scope.orderProp = 'age';

	
	$scope.gridOptions = { 
		data:  'phones',
		enablePinning: true,
        enableRowSelection: true,
        enableCellEdit: true,
		columnDefs: [{field:'id', displayName:'ID'}, {field:'name', displayName:'Name'}, {field:'age', displayName:'Age', groupable:true}],
		showGroupPanel: true,
		//filterOptions : {filterText: 'Test', useExternalFilter: false},
		showFilter : true,
		//jqueryUIDraggable: true,
	};
  }]);


/*
phonecatApp.controller('PhoneListCtrl', function ($scope, $http) {
  $http.get('phones/phones.json').
  //$http.get('test/test.json').
  //$http.get('http://127.0.0.1:28017/test/phones/').
    success(function(data) {
      //alert(data.d);
      $scope.phones = angular.fromJson(data.d);
    }).
    error(function(data) {
      alert(data);
    });
 
  $scope.orderProp = 'age';
});
*/

phonecatControllers.controller('PhoneDetailCtrl', ['$scope', '$routeParams', 'Phone',
  function($scope, $routeParams, Phone) {
    $scope.phone = Phone.get({phoneId: $routeParams.phoneId}, function(phone) {
      $scope.mainImageUrl = phone.images[0];
    });

    $scope.setImage = function(imageUrl) {
      $scope.mainImageUrl = imageUrl;
    }
  }]);
