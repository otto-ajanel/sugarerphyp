<?php

namespace App\Infrastructure\Persistence\Eloquent;

use App\Domain\Users\User;
use App\Domain\Users\UserRepository;
use App\Infrastructure\Persistence\Eloquent\Models\User as ModelsUser;

class UserRepositoryEloquent implements UserRepository{
    public function findByEmail(string $email): ?User
    {
        $dataUserModel = ModelsUser::where('email', $email)->first();
        return $dataUserModel ? null :null;

    } 
}