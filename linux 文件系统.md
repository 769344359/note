linux 的文件系统主要要介绍  
1 VFS  
2 file 指针  
3 inode结构  
3 具体的文件系统比如 ext2  
4 address_space  
```
struct address_space {
	struct inode		*host;		/* owner: inode, block_device */
	struct radix_tree_root	page_tree;	/* radix tree of all pages */
	spinlock_t		tree_lock;	/* and lock protecting it */
	atomic_t		i_mmap_writable;/* count VM_SHARED mappings */
	struct rb_root_cached	i_mmap;		/* tree of private and shared mappings */
	struct rw_semaphore	i_mmap_rwsem;	/* protect tree, count, list */
	/* Protected by tree_lock together with the radix tree */
	unsigned long		nrpages;	/* number of total pages */
	/* number of shadow or DAX exceptional entries */
	unsigned long		nrexceptional;
	pgoff_t			writeback_index;/* writeback starts here */
	const struct address_space_operations *a_ops;	/* methods */
	unsigned long		flags;		/* error bits */
	spinlock_t		private_lock;	/* for use by the address_space */
	gfp_t			gfp_mask;	/* implicit gfp mask for allocations */
	struct list_head	private_list;	/* for use by the address_space */
	void			*private_data;	/* ditto */
	errseq_t		wb_err;
} __attribute__((aligned(sizeof(long)))) __randomize_layout;
```

现在还有的疑问  
1 file   
2 页表  
3 inode   
4 page   
这四个的关系是怎么样的

---
我现在知道的是
vfs 是一层抽象层  
FILE * 是一个流,有缓冲,是用户态fd 的跟加深的封装  
fd  是进程的用户层的表示,可以通过fd 获取file  
file 属于进程  
inode 属于文件    
页表是用来使虚拟地址映射到相应的物理页    
