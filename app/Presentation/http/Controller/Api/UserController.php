<?php

namespace App\Presentation\http\Controller\Api;

use App\Application\Services\UserService;
use Swoole\Coroutine;
use Swoole\Coroutine\Channel;

class UserController
{
    private $userService;
        private $channel;

    public function __construct(UserService $userService)
    {
        $this->userService = $userService;
         $this->channel = new Channel(100);
    }

    public function index()
    {
        return $this->userService->getAllUsers();

}
}