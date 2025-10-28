<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Eloquent\Models;

use Hyperf\DbConnection\Model\Model;

/**
 * @property int $prodprec_id 
 * @property string $prodprec_precio 
 * @property string $prodprec_impuestopre 
 * @property int $product_id 
 */
class ProductsPrice extends Model
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'products_prices';

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
    protected array $casts = ['prodprec_id' => 'integer', 'product_id' => 'integer'];
}
