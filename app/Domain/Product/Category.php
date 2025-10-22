<?php
namespace App\Domain\Product;
class Category{
    
    public function __construct(
        public int $id_category, 
        public string $des_category, 
        )    {
    }   
}