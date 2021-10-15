import datetime

print(f"Normal fibnocci program")

series=10

def fibnocci(n):
	if n<=1:
		return 1
	else:
		return fibnocci(n-1)+fibnocci(n-2)

startTime=datetime.datetime.now()
result=[]
for i in range(series):
	result.append(fibnocci(i))

print("result - ",result)
print("time taken - ", datetime.datetime.now()-startTime)


print(f"Dynamic fibnocci program")
startTime=datetime.datetime.now()
result=[None for i in range(series)]

def fibnocci_dynamic(n):
	if result[n]!=None:
		return result[n]
	if n<=1:
		return 1
	else:
		return (fibnocci_dynamic(n-1)+fibnocci_dynamic(n-2))

for i in range(series):
	result[i]=fibnocci_dynamic(i)

print("result - ",result)
print("time taken - ", datetime.datetime.now()-startTime)