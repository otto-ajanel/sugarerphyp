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
use App\Presentation\http\Controller\Api\AuthController;
use App\Presentation\http\Controller\Api\UserController;
use Hyperf\HttpServer\Router\Router;


Router::post('/api/login', [AuthController::class, 'login']);
 
Router::get('/favicon.ico', function () {
    return '';
});

Router::addGroup("/api/v1", function(){
    
    Router::get("/user", function(){
        return ["messa"=>"Hello user with token"];
    });
},[
    'middleware'=>[AuthMiddlewareToken::class]
]);
