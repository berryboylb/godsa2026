### Key Insight
Duplicates can be detected either by sorting adjacent elements
or by tracking seen values using a hash map.

### Trade-offs
- Sorting avoids extra memory but mutates input
- Hash map preserves input and gives linear time

### Interview Choice
Hash map solution is preferred unless space is constrained.
