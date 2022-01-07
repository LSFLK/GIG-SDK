package config

/**
Set the GIG server API url here for crawlers
 */

//var ServerUrl = "https://api.gigdemo.opensource.lk"
var ServerUrl = "http://localhost"
var ApiUrl = ServerUrl + ":9000/api/"
var NERServerUrl = "http://18.221.69.238:8080/classify"
var NormalizeServer = ApiUrl + "normalize"
var OCRServer = ServerUrl + ":8081/extract?url="
