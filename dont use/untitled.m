

%% read in the availabilities
file_path = 'formatted_data.csv';
availability = readmatrix(file_path);
% disp(availability);


%% Constraints
% constraint 1: we can have uhhh max 5 meetings a day
% constraint 2: we can have max 3 meetings in a row
% constraint 3: we can have max uhhh 3 meetings per half day?
% goal 1: try to spread meetings across days
% goal 2: try to spread meetings equally across attendees
% goal 3: how do we set a trade-off for if someone can't make it to the meeting?

max_meetings_in_day = 5;
max_meetings_in_half_day = 3;
max_meetings_in_row = 3;



%% halfday_possibilities
% apply constraints 2 and 3 onto meetings

score = 0;

%% unsure how7 to approach the problem, fragmented it to:
% split to 4, 4, 4, 5, 4, 4, 4, 5


tic 
result = get_most_appts(1, availability, 1,1);
toc





function most_appts = get_most_appts(num,availability, starting, ending)
    half_day = 2^num-1;
    
    halfday_possibilities = [];
    for index = 0:half_day
    
        % constraint 3: if more than 3 meetings, get rid of the possibility
        val = dec2bin(index,num);
        if count(val,"1") <= 2
    
            % we convert it from string to int
            temp = val';
            temp2 = [];
            for x = 1:length(temp)
                temp2 = [temp2; str2double(temp(x))];
            end
            halfday_possibilities = [halfday_possibilities temp2];
        end
    end % now we have halfday_possibility, the entire set of feasible columns for half days


    for c1 = 1:length(halfday_possibilities) 
    for c2 = 1:length(halfday_possibilities) 
        for c3 = 1:length(halfday_possibilities) 
        for c4 = 1:length(halfday_possibilities) 
        for c5 = 1:length(halfday_possibilities) 
        for c6 = 1:length(halfday_possibilities) 
        for c7 = 1:length(halfday_possibilities) 
        for c8 = 1:length(halfday_possibilities) 
        for c9 = 1:length(halfday_possibilities) 
        for c10 = 1:length(halfday_possibilities) 


                for c11 = 1:length(halfday_possibilities) 
    for c12 = 1:length(halfday_possibilities) 
        for c13 = 1:length(halfday_possibilities) 
        for c14 = 1:length(halfday_possibilities) 
        for c15 = 1:length(halfday_possibilities) 
        for c16 = 1:length(halfday_possibilities) 
        for c17 = 1:length(halfday_possibilities) 
        for c18 = 1:length(halfday_possibilities) 
        for c19 = 1:length(halfday_possibilities) 
        for c20 = 1:length(halfday_possibilities) 


                for c21 = 1:length(halfday_possibilities) 
    for c22 = 1:length(halfday_possibilities) 
        for c23 = 1:length(halfday_possibilities) 
        for c24 = 1:length(halfday_possibilities) 
        for c25 = 1:length(halfday_possibilities) 
        for c26 = 1:length(halfday_possibilities) 
        for c27 = 1:length(halfday_possibilities) 
        for c28 = 1:length(halfday_possibilities) 
        for c29 = 1:length(halfday_possibilities) 
        for c30 = 1:length(halfday_possibilities) 

    for c31 = 1:length(halfday_possibilities) 
    for c32 = 1:length(halfday_possibilities) 
        for c33 = 1:length(halfday_possibilities) 
        for c34 = 1:length(halfday_possibilities) 
        for c35 = 1:length(halfday_possibilities) 
        for c36 = 1:length(halfday_possibilities) 
        for c37 = 1:length(halfday_possibilities) 
        for c38 = 1:length(halfday_possibilities) 
        for c39 = 1:length(halfday_possibilities) 
        for c40 = 1:length(halfday_possibilities) 

           possibility = halfday_possibilities(:,c1);
           availability(starting:ending, :);

       end
        end
               end
        end
               end
        end
               end
        end
               end
    end
           end
        end
               end
        end
               end
        end
               end
        end       
    end
                end       
        end
        end       
        end
        end       
        end
        end
            end
           end
        end
               end
        end
               end
        end
               end
        end       
    end
                end       
        end
        end       
        end


    most_appts;
end






