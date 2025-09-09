<?php

declare(strict_types=1);

namespace App\Infrastructure\Persistence\Eloquent\Models;

use Hyperf\Database\Model\Model;

/**
 * @property int $id_module 
 * @property string $module_name 
 * @property string $description 
 * @property string $icon_module 
 */
class Module extends TenantModel
{
    /**
     * The table associated with the model.
     */
    protected ?string $table = 'modules';

    /**
     * The attributes that are mass assignable.
     */
    protected array $fillable = [];

    /**
     * The attributes that should be cast to native types.
     */
    protected array $casts = ['id_module' => 'integer'];
}
