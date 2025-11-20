<?php

declare(strict_types=1);
/**
 * This file is part of Hyperf.
 *
 * @link     https://www.hyperf.io
 * @document https://hyperf.wiki
 * @contact  group@hyperf.io
 * @license  https://github.com/hyperf/hyperf/blob/master/LICENSE
 */

use App\Middleware\AuthMiddlewareToken;
use App\Presentation\http\Controllers\Api\AtributeController;
use App\Presentation\http\Controllers\Api\AuthController;
use App\Presentation\http\Controllers\Api\CategoryController;
use App\Presentation\http\Controllers\Api\ProductController;
use App\Presentation\http\Controllers\Api\UserController;
use Hyperf\HttpServer\Router\Router;


Router::post('/api/login', [AuthController::class, 'login']);
 
Router::get('/favicon.ico', function () {
    return '';
});

Router::addGroup("/api/v1", function(){
    
    Router::get("/user",[UserController::class, 'index']);
    Router::get("/permissionsbyuser",[UserController::class, 'permissionsByUser']);
    Router::get("/categories",[CategoryController::class, 'index']);
    Router::post("/createcategory",[CategoryController::class, 'store']);
    Router::post("/product",[ProductController::class, 'createProduct']);
    Router::get("/products",[ProductController::class, 'getAllProducts']);
    Router::get("/atributes",[AtributeController ::class, 'getAllAtributes']);
    Router::post("/createatribute",[AtributeController ::class, 'createAtribute']);
},[
    'middleware'=>[AuthMiddlewareToken::class]
]);