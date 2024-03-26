    # BE CAREFUL WHEN RUNNING THIS 	
    

    # hard code everything for the moment
    # this is very time consuming (30 minutes)

    # sort -u iter4_1_ab_1
    cat -n iter4_1_ab_1 | sort -uk2 | sort -nk1 | cut -f2- > _temp1_iter5_1
    # sort -u iter4_1_ae_1
    cat -n iter4_1_ae_1 | sort -uk2 | sort -nk1 | cut -f2- > _temp2_iter5_1
    # sort -u iter4_1_ad_1
    cat -n iter4_1_ad_1 | sort -uk2 | sort -nk1 | cut -f2- > _temp3_iter5_1
    # sort -u iter4_1_ac_1
    cat -n iter4_1_ac_1 | sort -uk2 | sort -nk1 | cut -f2- > _temp4_iter5_1


    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp1_iter5_1" "_temp2_iter5_1" > _temp5_iter5_1
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp3_iter5_1" "_temp4_iter5_1" > _temp5_iter6_1
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp5_iter5_1" "_temp5_iter6_1" > iter5_1
#     # 

    # sort -u iter4_1_bc_1
    cat -n iter4_1_bc_1 | sort -uk2 | sort -nk1 | cut -f2- > _temp1_iter5_2
    # sort -u iter4_1_ab_2
    cat -n iter4_1_ab_2 | sort -uk2 | sort -nk1 | cut -f2- > _temp2_iter5_2
    # sort -u iter4_1_bd_1
    cat -n iter4_1_bd_1 | sort -uk2 | sort -nk1 | cut -f2- > _temp3_iter5_2
    # sort -u iter4_1_be_1 
    cat -n iter4_1_be_1 | sort -uk2 | sort -nk1 | cut -f2- > _temp4_iter5_2

    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp1_iter5_2" "_temp2_iter5_2" > _temp5_iter5_2
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp3_iter5_2" "_temp4_iter5_2" > _temp5_iter6_2
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp5_iter5_2" "_temp5_iter6_2" > iter5_2


# # 

    # sort -u iter4_1_cd_1
    cat -n iter4_1_cd_1 |  sort -uk2 | sort -nk1 | cut -f2- > _temp1_iter5_3
    # sort -u iter4_1_ac_2
    cat -n iter4_1_ac_2 |  sort -uk2 | sort -nk1 | cut -f2- > _temp2_iter5_3
    # sort -u iter4_1_ce_1
    cat -n iter4_1_ce_1 |  sort -uk2 | sort -nk1 | cut -f2- > _temp3_iter5_3
    # sort -u iter4_1_bc_2
    cat -n iter4_1_bc_2 |  sort -uk2 | sort -nk1 | cut -f2- > _temp4_iter5_3


    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp1_iter5_3" "_temp2_iter5_3" > _temp5_iter5_3
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp3_iter5_3" "_temp4_iter5_3" > _temp5_iter6_3
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp5_iter5_3" "_temp5_iter6_3" > iter5_3

#     # 

    # sort -u iter4_1_ad_2
    cat -n iter4_1_ad_2 | sort -uk2 | sort -nk1 | cut -f2- > _temp1_iter5_4
    # sort -u iter4_1_cd_2
    cat -n iter4_1_cd_2 | sort -uk2 | sort -nk1 | cut -f2- > _temp2_iter5_4
    # sort -u iter4_1_bd_2
    cat -n iter4_1_bd_2 | sort -uk2 | sort -nk1 | cut -f2- > _temp3_iter5_4
    # sort -u iter4_1_de_1
    cat -n iter4_1_de_1 | sort -uk2 | sort -nk1 | cut -f2- > _temp4_iter5_4

    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp1_iter5_4" "_temp2_iter5_4" > _temp5_iter5_4
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp3_iter5_4" "_temp4_iter5_4" > _temp5_iter6_4
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp5_iter5_4" "_temp5_iter6_4" > iter5_4

#     # 

    # sort -u iter4_1_ce_2
    cat -n iter4_1_ce_2 | sort -uk2 | sort -nk1 | cut -f2- > _temp1_iter5_5
    # sort -u iter4_1_ae_2
    cat -n iter4_1_ae_2 | sort -uk2 | sort -nk1 | cut -f2- > _temp2_iter5_5
    # sort -u iter4_1_be_2
    cat -n iter4_1_be_2 | sort -uk2 | sort -nk1 | cut -f2- > _temp3_iter5_5
    # sort -u iter4_1_de_2 
    cat -n iter4_1_de_2  | sort -uk2 | sort -nk1 | cut -f2- > _temp4_iter5_5


    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp1_iter5_5" "_temp2_iter5_5" > _temp5_iter5_5
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp3_iter5_5" "_temp4_iter5_5" > _temp5_iter6_5
    awk 'NR==FNR { lines[$0]=1; next } $0 in lines' "_temp5_iter5_5" "_temp5_iter6_5" > iter5_5

# 


    # comm -12 <(sort iter4_1_ab_1) <(sort iter4_1_ac_1) > _temp1_iter5_1
    # comm -12 <(sort iter4_1_ab_1) <(sort iter4_1_ac_1) > _temp1_iter5_2

















    # # this gives the viable columns of b 
    # awk 'NR==FNR { lines[$0]=1; next } $0 in lines' iter4_1_ab_2 iter4_1_bc_1 > b_intersection

    # # this gives the viable columns of c
    # awk 'NR==FNR { lines[$0]=1; next } $0 in lines' iter4_1_bc_2 iter4_1_cd_1 > c_intersection

    # # this gives the viable columns of d
    # awk 'NR==FNR { lines[$0]=1; next } $0 in lines' iter4_1_cd_2 iter4_1_de_1 > d_intersection


    # # https://unix.stackexchange.com/questions/444795/remove-duplicate-lines-from-a-file-but-leave-1-occurrence
	# sort -u b_intersection
	# cat -n b_intersection | sort -uk2 | sort -nk1 | cut -f2-


    # # https://unix.stackexchange.com/questions/444795/remove-duplicate-lines-from-a-file-but-leave-1-occurrence
	# sort -u c_intersection
	# cat -n c_intersection | sort -uk2 | sort -nk1 | cut -f2-


    # # https://unix.stackexchange.com/questions/444795/remove-duplicate-lines-from-a-file-but-leave-1-occurrence
	# sort -u d_intersection
	# cat -n d_intersection | sort -uk2 | sort -nk1 | cut -f2-