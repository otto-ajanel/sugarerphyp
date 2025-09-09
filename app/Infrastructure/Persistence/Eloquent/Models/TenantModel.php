<?php

namespace App\Infrastructure\Persistence\Eloquent\Models;

use Hyperf\Context\Context;
use Hyperf\DbConnection\Model\Model;

class TenantModel extends Model
{
     protected ?string $connection = null;

    public function setConnectionName(string $pool): static
    {
        $this->connection = $pool;
        return $this;
    }

    public function getConnectionName(): ?string
    {
        // Si $connection ya está seteado, úsalo
        return $this->connection ?? Context::get('tenant');
    }
      
}