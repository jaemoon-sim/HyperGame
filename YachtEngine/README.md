## build
```
go build --o yachtengine main.go
```

## run
```
./yachtengine --output log.json asdf,9910 zxcv,9990
```

### params
|key|required|설명|예시|
|---|---|---|---|
|output|N|yachtengine 실행 후 전체 로그 출력 path. 비어있을 경우 stdout|`--output ./log.json`|
|(없음)|Y|`이름,port번호` 형식. 여러 명 참여 가능. 이름 중복 가능|`jb,1234 jb2,1234 jb3,1234`|

## output

|key|설명|
|---|---
|log|player의 decision이 발생하기 직전의 state와 player의 decision을 묶어서 json array로 가지고 있음
|`log > state`|player의 decision이 발생하기 직전의 state
|`log > decision`| `log > state`에 대한 player의 decision
|`final`| 다 끝난 뒤 게임 state|

```
{
 "players": [
  "jb",
  "jb2"
 ],
 "log": [
  {
   "state": {
    "turn": 1,
    "player": "jb",
    "trial": 1,
    "dices": [
     2,
     6,
     2,
     2,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": null,
      "sixes": null,
      "subtotal": 0,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 0
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": null,
      "sixes": null,
      "subtotal": 0,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 0
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "sixes"
   }
  },
  {
   "state": {
    "turn": 1,
    "player": "jb2",
    "trial": 1,
    "dices": [
     5,
     2,
     6,
     4,
     6
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": null,
      "sixes": 6,
      "subtotal": 6,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 6
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": null,
      "sixes": null,
      "subtotal": 0,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 0
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "sixes"
   }
  },
  {
   "state": {
    "turn": 2,
    "player": "jb",
    "trial": 1,
    "dices": [
     2,
     4,
     2,
     6,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": null,
      "sixes": 6,
      "subtotal": 6,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 6
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": null,
      "sixes": 12,
      "subtotal": 12,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 12
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fives"
   }
  },
  {
   "state": {
    "turn": 2,
    "player": "jb2",
    "trial": 1,
    "dices": [
     5,
     2,
     2,
     6,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": 5,
      "sixes": 6,
      "subtotal": 11,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 11
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": null,
      "sixes": 12,
      "subtotal": 12,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 12
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fives"
   }
  },
  {
   "state": {
    "turn": 3,
    "player": "jb",
    "trial": 1,
    "dices": [
     5,
     1,
     1,
     4,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": 5,
      "sixes": 6,
      "subtotal": 11,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 11
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": 5,
      "sixes": 12,
      "subtotal": 17,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 17
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fours"
   }
  },
  {
   "state": {
    "turn": 3,
    "player": "jb2",
    "trial": 1,
    "dices": [
     1,
     4,
     5,
     2,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 19,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 19
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": null,
      "fives": 5,
      "sixes": 12,
      "subtotal": 17,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 17
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fours"
   }
  },
  {
   "state": {
    "turn": 4,
    "player": "jb",
    "trial": 1,
    "dices": [
     1,
     5,
     4,
     4,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 19,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 19
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 25,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 25
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "choice"
   }
  },
  {
   "state": {
    "turn": 4,
    "player": "jb2",
    "trial": 1,
    "dices": [
     2,
     5,
     2,
     1,
     6
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 19,
      "choice": 19,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 38
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 25,
      "choice": null,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 25
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "choice"
   }
  },
  {
   "state": {
    "turn": 5,
    "player": "jb",
    "trial": 1,
    "dices": [
     1,
     1,
     5,
     3,
     3
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 19,
      "choice": 19,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 38
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 25,
      "choice": 16,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 41
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "threes"
   }
  },
  {
   "state": {
    "turn": 5,
    "player": "jb2",
    "trial": 1,
    "dices": [
     6,
     1,
     6,
     1,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 25,
      "choice": 19,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 44
     },
     "jb2": {
      "aces": null,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 25,
      "choice": 16,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 41
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "aces"
   }
  },
  {
   "state": {
    "turn": 6,
    "player": "jb",
    "trial": 1,
    "dices": [
     6,
     5,
     2,
     1,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": null,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 25,
      "choice": 19,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 44
     },
     "jb2": {
      "aces": 2,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 27,
      "choice": 16,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 43
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "deuces"
   }
  },
  {
   "state": {
    "turn": 6,
    "player": "jb2",
    "trial": 1,
    "dices": [
     6,
     4,
     5,
     4,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 27,
      "choice": 19,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 46
     },
     "jb2": {
      "aces": 2,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 27,
      "choice": 16,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 43
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fullHouse"
   }
  },
  {
   "state": {
    "turn": 7,
    "player": "jb",
    "trial": 1,
    "dices": [
     3,
     6,
     2,
     6,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 27,
      "choice": 19,
      "fourKind": null,
      "fullHouse": null,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 46
     },
     "jb2": {
      "aces": 2,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 27,
      "choice": 16,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 43
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fullHouse"
   }
  },
  {
   "state": {
    "turn": 7,
    "player": "jb2",
    "trial": 1,
    "dices": [
     4,
     2,
     6,
     2,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 27,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 46
     },
     "jb2": {
      "aces": 2,
      "deuces": null,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 27,
      "choice": 16,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 43
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "deuces"
   }
  },
  {
   "state": {
    "turn": 8,
    "player": "jb",
    "trial": 1,
    "dices": [
     1,
     3,
     4,
     3,
     3
    ],
    "scoreBoard": {
     "jb": {
      "aces": null,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 27,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 46
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 33,
      "choice": 16,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 49
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "aces"
   }
  },
  {
   "state": {
    "turn": 8,
    "player": "jb2",
    "trial": 1,
    "dices": [
     1,
     3,
     3,
     3,
     3
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 47
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 33,
      "choice": 16,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 49
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fourKind"
   }
  },
  {
   "state": {
    "turn": 9,
    "player": "jb",
    "trial": 1,
    "dices": [
     6,
     6,
     6,
     3,
     6
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": null,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 47
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 33,
      "choice": 16,
      "fourKind": 13,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 62
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "fourKind"
   }
  },
  {
   "state": {
    "turn": 9,
    "player": "jb2",
    "trial": 1,
    "dices": [
     6,
     4,
     3,
     5,
     1
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": 27,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 74
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": null,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 33,
      "choice": 16,
      "fourKind": 13,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 62
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "threes"
   }
  },
  {
   "state": {
    "turn": 10,
    "player": "jb",
    "trial": 1,
    "dices": [
     5,
     4,
     2,
     5,
     4
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": 27,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 74
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 36,
      "choice": 16,
      "fourKind": 13,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 65
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "largeStraight"
   }
  },
  {
   "state": {
    "turn": 10,
    "player": "jb2",
    "trial": 1,
    "dices": [
     3,
     1,
     4,
     5,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": 27,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": 0,
      "yacht": null,
      "total": 74
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 36,
      "choice": 16,
      "fourKind": 13,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": null,
      "yacht": null,
      "total": 65
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "smallStraight"
   }
  },
  {
   "state": {
    "turn": 11,
    "player": "jb",
    "trial": 1,
    "dices": [
     2,
     2,
     4,
     6,
     1
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": 27,
      "fullHouse": 0,
      "smallStraight": null,
      "largeStraight": 0,
      "yacht": null,
      "total": 74
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 36,
      "choice": 16,
      "fourKind": 13,
      "fullHouse": 0,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null,
      "total": 80
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "smallStraight"
   }
  },
  {
   "state": {
    "turn": 11,
    "player": "jb2",
    "trial": 1,
    "dices": [
     5,
     5,
     2,
     6,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": 27,
      "fullHouse": 0,
      "smallStraight": 0,
      "largeStraight": 0,
      "yacht": null,
      "total": 74
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 36,
      "choice": 16,
      "fourKind": 13,
      "fullHouse": 0,
      "smallStraight": 15,
      "largeStraight": null,
      "yacht": null,
      "total": 80
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "largeStraight"
   }
  },
  {
   "state": {
    "turn": 12,
    "player": "jb",
    "trial": 1,
    "dices": [
     3,
     4,
     4,
     3,
     2
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": 27,
      "fullHouse": 0,
      "smallStraight": 0,
      "largeStraight": 0,
      "yacht": null,
      "total": 74
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 36,
      "choice": 16,
      "fourKind": 13,
      "fullHouse": 0,
      "smallStraight": 15,
      "largeStraight": 0,
      "yacht": null,
      "total": 80
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "yacht"
   }
  },
  {
   "state": {
    "turn": 12,
    "player": "jb2",
    "trial": 1,
    "dices": [
     2,
     5,
     6,
     4,
     5
    ],
    "scoreBoard": {
     "jb": {
      "aces": 1,
      "deuces": 2,
      "threes": 6,
      "fours": 8,
      "fives": 5,
      "sixes": 6,
      "subtotal": 28,
      "choice": 19,
      "fourKind": 27,
      "fullHouse": 0,
      "smallStraight": 0,
      "largeStraight": 0,
      "yacht": 0,
      "total": 74
     },
     "jb2": {
      "aces": 2,
      "deuces": 6,
      "threes": 3,
      "fours": 8,
      "fives": 5,
      "sixes": 12,
      "subtotal": 36,
      "choice": 16,
      "fourKind": 13,
      "fullHouse": 0,
      "smallStraight": 15,
      "largeStraight": 0,
      "yacht": null,
      "total": 80
     }
    }
   },
   "decision": {
    "keep": [
     0,
     1,
     2,
     3,
     4
    ],
    "choice": "yacht"
   }
  }
 ],
 "final": {
  "dices": [
   0,
   0,
   0,
   0,
   0
  ],
  "scoreBoard": {
   "jb": {
    "aces": 1,
    "deuces": 2,
    "threes": 6,
    "fours": 8,
    "fives": 5,
    "sixes": 6,
    "subtotal": 28,
    "choice": 19,
    "fourKind": 27,
    "fullHouse": 0,
    "smallStraight": 0,
    "largeStraight": 0,
    "yacht": 0,
    "total": 74
   },
   "jb2": {
    "aces": 2,
    "deuces": 6,
    "threes": 3,
    "fours": 8,
    "fives": 5,
    "sixes": 12,
    "subtotal": 36,
    "choice": 16,
    "fourKind": 13,
    "fullHouse": 0,
    "smallStraight": 15,
    "largeStraight": 0,
    "yacht": 0,
    "total": 80
   }
  }
 }
}
```
