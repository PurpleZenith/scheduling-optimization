    # BE CAREFUL WHEN RUNNING THIS 	
    
    # this gives the viable columns of b 
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' iter4_1_ab_2 iter4_1_bc_1 > b_intersection

    # this gives the viable columns of c
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' iter4_1_bc_2 iter4_1_cd_1 > c_intersection

    # this gives the viable columns of d
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' iter4_1_cd_2 iter4_1_de_1 > d_intersection


    # https://unix.stackexchange.com/questions/444795/remove-duplicate-lines-from-a-file-but-leave-1-occurrence
	sort -u b_intersection
	cat -n b_intersection | sort -uk2 | sort -nk1 | cut -f2-


    # https://unix.stackexchange.com/questions/444795/remove-duplicate-lines-from-a-file-but-leave-1-occurrence
	sort -u c_intersection
	cat -n c_intersection | sort -uk2 | sort -nk1 | cut -f2-


    # https://unix.stackexchange.com/questions/444795/remove-duplicate-lines-from-a-file-but-leave-1-occurrence
	sort -u d_intersection
	cat -n d_intersection | sort -uk2 | sort -nk1 | cut -f2-