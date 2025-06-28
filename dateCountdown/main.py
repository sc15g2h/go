from datetime import datetime
import time

import re


def valiDate(date):
    try:
        bool(datetime.strptime(dateinput, '%d-%m-%Y'))
        return True
    except: 
        print("\n* * * * * * * * * * * * Error * * * * * * * * * * * *\nIncorrect date format entered.\nPlease try again using format DD-MM-YYYY\n* * * * * * * * * * * * * * * * * * * * * * * * * * *\n")
        return False

def inPast(days):
    if days < 0:
        print("\n* * * * * * * * * * * * Error * * * * * * * * * * * *\nThe date you have entered is in the past\n* * * * * * * * * * * * * * * * * * * * * * * * * * *\n")
        return False
    elif days == 0:
        print("Thats today!")
        return False
    else: 
        return True


def difference(today, date):
    dt1 = time.mktime(time.strptime(today, '%d-%m-%Y'))
    dt2 = time.mktime(time.strptime(date, '%d-%m-%Y'))
    delta = dt2-dt1
    diff = int(delta / 86400)
    
    return diff


while True:
    print("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -\n")
    print("Choose a date and I'll tell you how many days you have to wait til that day! \n*Date should be in format DD-MM-YYYY")
    dateinput = input().strip()

    # exit prog if user requests 
    if(dateinput.lower() == "exit"):
        break


    today = datetime.today().strftime('%d-%m-%Y')

    if(valiDate(dateinput)):
        d = difference(today, dateinput)
        if(inPast(d)):
            # Todays date
            print("You have " + str(d) + " days left!\n")
    
        


    



