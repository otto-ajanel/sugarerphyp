<?php

namespace App\Presentation\http\Controllers\Api;

use App\Application\Services\UserService;
use Hyperf\Context\Context;
use Swoole\Coroutine;
use Swoole\Coroutine\Channel;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface;
use Psr\Http\Message\ResponseInterface as PsrRes;

class UserController
{
    private $userService;
    private $channel;
    private RequestInterface $req;
    private ResponseInterface $res;

    

    public function __construct(
        UserService $userService,
        RequestInterface $request,
        ResponseInterface $response
        )
    {
        $this->userService = $userService;
        $this->req = $request;
        $this->res = $response;
    }

    public function  index() :PsrRes
    {
        try {
            $reqData = $this->req->all();

            $data=$this->userService->getAllUsers($reqData);
            return $this->res->json($data);

        } catch (\Throwable $th) {
            return $this->res->json(['error' => $th->getMessage()])->withStatus(501);
        }
    }

    function permissionsByUser(){
        $userData = Context::get('userData');
        $data = $this->userService->getPermissionsByUser($userData['id_user']);

        return $data;
    }

}
