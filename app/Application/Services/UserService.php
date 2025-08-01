<?php

namespace App\Application\Services;

use App\Infrastructure\Persistence\Eloquent\Models\User;

class UserService {
    private Authservice $authService;
    public function __construct(
        Authservice $authservice
    )
    {
        $this->authService = $authservice;
    }

    
    public function loginService($email, $password){
       $userData = User::where([
            'email' => $email,
            'password' =>$password,
            'active'=> true
        ])
        ->first(); 
        
        if($userData==null){

            return null;
        }

        $token = $this->authService->encriptData($userData);
        
        $userData->token= $token;
        
        return $userData;   

    }

}