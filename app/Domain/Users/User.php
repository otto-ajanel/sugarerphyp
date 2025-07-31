<?php

namespace App\Domain\Users;
class User{
    
    public function __construct(
        public int $id_user, 
        public string $username, 
        public string $email, 
        public string $password, 
        public int $id_usertype, 
        public string $created_at, 
        public int $active, 
        public int $id_tenance,  
        )    {
    }   
}