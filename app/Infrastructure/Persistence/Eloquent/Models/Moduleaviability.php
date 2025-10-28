<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Eloquent\Models;

use Hyperf\DbConnection\Model\Model;

/**
 * @property int $modavia_id 
 * @property int $module_id 
 * @property int $product_id 
 */
class Moduleaviability extends Model
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'moduleaviability';

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
    protected array $casts = ['modavia_id' => 'integer', 'module_id' => 'integer', 'product_id' => 'integer'];
}
