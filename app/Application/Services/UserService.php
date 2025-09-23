<?php

namespace App\Application\Services;

use App\Infrastructure\Persistence\Eloquent\Models\Module;
use App\Infrastructure\Persistence\Eloquent\Models\User;
use Hyperf\Context\Context;
use Hyperf\DbConnection\Db;

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
        ->join('tenants', 'tenants.id_tenant', '=', 'users.id_tenant')
        ->first(); 
        
        if($userData==null){

            return null;
        }

        $token = $this->authService->encriptData($userData);
        
        $userData->token= $token;
        
        return $userData;   

    }

    public function getAllUsers($reqData){
        $perPage= $reqData['perPage'] ?? 10;
        $noPage = $reqData['noPage'] ?? 1;
        $dataUsers= User::query()
        ->paginate($perPage, ['*'], 'page', $noPage);
        return $dataUsers;
    }

    public function getPermissionsByUser($userId){

       return  Module::select("*")
        ->join("menus", "menus.id_module", "modules.id_module")
        ->join("userpermissions","menus.id_menu","=","userpermissions.id_menu")
        ->where("userpermissions.id_user", $userId)
        ->get(); 
    }
}