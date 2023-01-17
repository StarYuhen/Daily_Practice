<?php
namespace Driver;

/**
 * @Table(name="users", database="market")
 */
class UsersEntity
{

    /**
     * @Column(name="uid", type="string", length="10", isPrimary=1)
     */
    public $uid;


    /**
     * @Column(name="uname", type="string", length="8")
     */
    public $uname;


    /**
     * @Column(name="password", type="string", length="8")
     */
    public $password;

	
	
}