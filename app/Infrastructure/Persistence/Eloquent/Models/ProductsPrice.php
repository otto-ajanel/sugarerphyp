<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Eloquent\Models;

use Hyperf\DbConnection\Model\Model;

/**
 * @property int $prodprec_id 
 * @property string $prodprec_price 
 * @property string $prodprec_taxprice 
 * @property int $product_id
 * @property string $prodprec_purchase 
 * @property string $producprec_purchasetax  
 */
class ProductsPrice extends TenantModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'products_prices';

    /**
     * The connection name for the model.
     */

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = [
        'prodprec_price',
        'prodprec_taxprice',
        'product_id',
        'prodprec_purchase',
        'producprec_purchasetax'    
    ];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['prodprec_id' => 'integer', 'product_id' => 'integer'];
    protected string $primaryKey = 'prodprec_id';
    public bool $timestamps = false;
}
