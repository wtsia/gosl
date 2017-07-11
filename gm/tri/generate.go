// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tri

// TODO: move Tetrahedra to tet package

/* Definitions
 *
 *   2D:
 *              Nodes                 Edges
 *
 *    y           2
 *    |           @                     @
 *    +--x       / \                   / \
 *            5 /   \ 4               /   \
 *             @     @             2 /     \ 1
 *            /       \             /       \
 *           /         \           /         \
 *          @-----@-----@         @-----------@
 *         0      3      1              0
 *
 *
 *   3D:              Nodes                                     Faces
 *     z
 *     |                |                                          |
 *    ,+--y             3                                          +
 *  x'                 /|`.                                       /|`.
 *                     ||  `,                                     ||  `,
 *                    / |    ',                                  / |    ',
 *                    | |      \                                 | |      \
 *                   /  |       `.                              /  |       `.
 *                   |  |         `,                            |  |         `,
 *                  /   7            9                         /   |           ',
 *                  |   |             \                        |   |    --.      \
 *                 /    |              `.                     /    |    |0|       `.
 *                 |    |                ',                   | /| |    ``-         ',
 *                8     |                  \                 / |1| |                  \
 *                |     0 ,,_               `.               | |/  + ,,_               `.
 *               |     /     ``'-., 6         `.            |     /     ``'-.,,.         `.
 *               |    /               `''-.,,_  ',          |    /     --       ``''-.,,_  ',
 *              |    /                        ``'2 ,,      |    /      \3\               ``'+ ,,
 *              |   '                       ,.-``          |  ,'        ``             ,.-``
 *             |   4                   _,-'`              |  /     ..             _,-'`
 *             ' /                 ,.'`                   ' /     /2/         ,.'`
 *            | /             _ 5 `                      | /      ''     _.''`
 *            '/          ,-'`                           '/          ,-'`
 *           |/      ,.-``                              |/      ,.-``
 *           /  _,-``                                   /  _,-``
 *          1 '`                                       +.'`
 *
 *     3D:
 *          Edges
 *
 *                |
 *                +
 *               /|`.
 *               ||  `,
 *              / |    ',
 *              | |      \
 *             /  |       `.
 *             |  |         `,5
 *            /   |3          ',
 *            |   |             \
 *           /    |              `.
 *         4 |    |                ',
 *          /     |                  \
 *          |     +.,,_               `.
 *         |     /     ``'-.,,.         `.
 *         |    /          2   ``''-.,,_  ',
 *        |    /                        ``'+ ,,
 *        |  ,'                       ,.-``
 *       |  / 0                  _,-'`
 *       ' /                 ,.'`
 *      | /             _.''`
 *      '/          ,-'`   1
 *     |/      ,.-``
 *     /  _,-``
 *    +.'`
 *
 */
