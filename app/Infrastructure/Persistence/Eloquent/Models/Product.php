<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Eloquent\Models;

use Hyperf\DbConnection\Model\Model;

/**
 * @property int $product_id 
 * @property string $product_sku 
 * @property string $product_des 
 * @property int $category_id 
 * @property int $typeproduct_id 
 * @property int $billingpolicy_id 
 * @property int $status_id 
 */
class Product extends TenantModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'products';



    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = [
        'product_sku',
        'product_des',
        'category_id',
        'typeproduct_id',
        'billingpolicy_id',
        'status_id'
    ];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['product_id' => 'integer', 'category_id' => 'integer', 'typeproduct_id' => 'integer', 'billingpolicy_id' => 'integer', 'product_status' => 'integer'];
    protected string $primaryKey = 'product_id';
    public bool $timestamps = false;
}
