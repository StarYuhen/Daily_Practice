<?php
namespace Driver;

/**
 * @Table(name="type", database="market")
 */
class TypeEntity
{

    /**
     * @Column(name="tid", type="string", length="10", isPrimary=1)
     */
    public $tid;


    /**
     * @Column(name="tname", type="string", length="15")
     */
    public $tname;


    /**
     * @Column(name="note", type="text")
     */
    public $note;

	
	
}