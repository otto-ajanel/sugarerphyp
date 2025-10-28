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
 * @property int $product_status 
 */
class Product extends Model
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'products';

    /**
     * The connection name for the model.
     */
    protected ?string $connection = 'pgsql';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = [];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['product_id' => 'integer', 'category_id' => 'integer', 'typeproduct_id' => 'integer', 'billingpolicy_id' => 'integer', 'product_status' => 'integer'];
}
