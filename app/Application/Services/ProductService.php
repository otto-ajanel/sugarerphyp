<?php

namespace App\Application\Services;

use App\Infrastructure\Persistence\Eloquent\Models\Moduleaviability;
use App\Infrastructure\Persistence\Eloquent\Models\Product;
use App\Infrastructure\Persistence\Eloquent\Models\ProductsPrice;

class ProductService
{
    public function createProduct(array $data): array
    {

    $newproduct = Product::create([
        'product_sku' => $data['barcode'],
        'product_des' => $data['nameProduct'],
        'status_id' => 1,
        'category_id' => $data['categoryId']['code'],
        'typeproduct_id' => $data['typeproduct'],
        'billingpolicy_id' => $data['billingPolicy']['code']
        ]
    );
    if ($newproduct) {

        ProductsPrice::create([
            'prodprec_price' => $data['salePrice'],
            'prodprec_taxprice' => $data['taxsale'],
            'product_id' => $newproduct->product_id,
            'prodprec_purchase' => $data['costPrice'],
            'producprec_purchasetax' => $data['taxpurchase']
        ]);
        
       if ($data['enableFor']['sale']) {
        # code...
        Moduleaviability::create([
            'module_id' => 12,
            'product_id' => $newproduct->product_id
        ]);
    }
    if ($data['enableFor']['pos']) {
        # code...
        Moduleaviability::create([
            'module_id' => 26,
            'product_id' => $newproduct->product_id
        ]);
    }
        
    
      if ($data['enableFor']['purchase']) {
        # code...
        Moduleaviability::create([
            'module_id' => 5,
            'product_id' => $newproduct->product_id
        ]);
    }

    }

    return
        $newproduct->toArray();
    }

    public function getAllProducts()
    {
        return Product::all()->toArray();
    }
}