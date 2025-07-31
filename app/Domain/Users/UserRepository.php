<?php
namespace App\Domain\Users;


interface UserRepository
{
    public  function findByEmail(string $email): ?User;
}