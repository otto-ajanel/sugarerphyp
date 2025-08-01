<?php

declare(strict_types=1);

namespace App\Presentation\http\Controller\Api;

use App\Application\Services\UserService;
use Hyperf\DbConnection\Db;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface;
use Psr\Http\Message\ResponseInterface as PsrRes;

class AuthController
{

    private UserService $userService;
    private RequestInterface $req;
    private ResponseInterface $res;

    public function __construct(
        UserService $userService,
        RequestInterface $request,
        ResponseInterface $response
    ) {
        $this->userService = $userService;
        $this->req = $request;
        $this->res = $response;
    }

    public function login(): PsrRes
    {
        try {
            $reqData = $this->req->all();
            $dataUser  = $this->userService->loginService($reqData['email'], $reqData['password']);
            if ($dataUser == null) {
                return $this->res->json(["Message" => "Not found user"])->withStatus(404);
            }

            return $this->res->json($dataUser);
         } catch (\Throwable $th) {
            return $this->res->json(["message" => $th])->withStatus(501);
        } 
    }
}
