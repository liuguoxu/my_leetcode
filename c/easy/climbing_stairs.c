#include <stdio.h>

int climb(int n)
{
	if(n == 1) {
		return 1;
	}else if(n == 2) {
		return 2;
	}

	return climb(n-1) + climb(n-2);
}

void main()
{
	printf("%d\n", climb(40));
}
