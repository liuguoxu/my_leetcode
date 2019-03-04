//
// Created by jere on 2019-3-4.
//

/*
 *You are given two non-empty linked lists representing two non-negative integers.
 * The digits are stored in reverse order and each of their nodes contain a single digit.
 * Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */

/**
 * Definition for singly-linked list.
 */
struct ListNode {
    int val;
    struct ListNode *next;
};

struct ListNode *addTwoNumbers(struct ListNode *l1, struct ListNode *l2)
{
    struct ListNode *ret = malloc(sizeof(struct ListNode));
    ret->next = NULL;
    struct ListNode *cur = ret;

    int m = 0;
    while (l1 || l2) {
        cur->next = malloc(sizeof(struct ListNode));
        cur->next->next = NULL;
        cur = cur->next;

        int sum = (l1 ? l1->val : 0) + (l2 ? l2->val : 0) + m;
        m = sum / 10;

        cur->val = sum % 10;

        if (l1) l1 = l1->next;
        if (l2) l2 = l2->next;
    }

    if(m) {
        cur->next = malloc(sizeof(struct ListNode));
        cur = cur->next;
        cur->val = 1;
        cur->next = NULL;
    }

    return ret->next;
}