<?php
namespace App\Domain\Product;
class Product{
    public function __construct(
        public int $id_product, 
        public string $des_product, 
        public int $id_category, 
        public float $price, 
        )    {
    }   
}
