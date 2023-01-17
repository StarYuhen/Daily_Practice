<?php
namespace Driver;

/**
 * @Table(name="goods", database="market")
 */
class GoodsEntity
{

    /**
     * @Column(name="gid", type="string", length="10", isPrimary=1)
     */
    public $gid;


    /**
     * @Column(name="gname", type="string", length="8")
     */
    public $gname;


    /**
     * @Column(name="price", type="float")
     */
    public $price;


    /**
     * @Column(name="count", type="int")
     */
    public $count;


    /**
     * @Column(name="on1", type="string", length="1")
     */
    public $on1;


    /**
     * @Column(name="recommend", type="string", length="1")
     */
    public $recommend;


    /**
     * @Column(name="type", type="string", length="10")
     */
    public $type;

	
	
}